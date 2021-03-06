/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package broker

import (
	"github.com/apache/incubator-servicecomb-service-center/server/core/backend/store"
)

var (
	PARTICIPANT  store.StoreType
	VERSION      store.StoreType
	PACT         store.StoreType
	PACT_VERSION store.StoreType
	PACT_TAG     store.StoreType
	VERIFICATION store.StoreType
	PACT_LATEST  store.StoreType
)

var brokerKvStore = &BKvStore{}

func init() {
	PARTICIPANT = store.Store().MustInstall(store.NewEntity("PARTICIPANT", GetBrokerParticipantKey("")))
	VERSION = store.Store().MustInstall(store.NewEntity("VERSION", GetBrokerVersionKey("")))
	PACT = store.Store().MustInstall(store.NewEntity("PACT", GetBrokerPactKey("")))
	PACT_VERSION = store.Store().MustInstall(store.NewEntity("PACT_VERSION", GetBrokerPactVersionKey("")))
	PACT_TAG = store.Store().MustInstall(store.NewEntity("PACT_TAG", GetBrokerTagKey("")))
	VERIFICATION = store.Store().MustInstall(store.NewEntity("VERIFICATION", GetBrokerVerificationKey("")))
	PACT_LATEST = store.Store().MustInstall(store.NewEntity("PACT_LATEST", GetBrokerLatestKey("")))

}

type BKvStore struct {
}

func (s *BKvStore) Participant() *store.Indexer {
	return store.Store().Entity(PARTICIPANT)
}

func (s *BKvStore) Version() *store.Indexer {
	return store.Store().Entity(VERSION)
}

func (s *BKvStore) Pact() *store.Indexer {
	return store.Store().Entity(PACT)
}

func (s *BKvStore) PactVersion() *store.Indexer {
	return store.Store().Entity(PACT_VERSION)
}

func (s *BKvStore) PactTag() *store.Indexer {
	return store.Store().Entity(PACT_TAG)
}

func (s *BKvStore) Verification() *store.Indexer {
	return store.Store().Entity(VERIFICATION)
}

func (s *BKvStore) PactLatest() *store.Indexer {
	return store.Store().Entity(PACT_LATEST)
}

func Store() *BKvStore {
	return brokerKvStore
}
