// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchStartStreamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *BatchStartStreamsRequest
	GetId() *string
	SetOwnerId(v int64) *BatchStartStreamsRequest
	GetOwnerId() *int64
}

type BatchStartStreamsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 323*****997-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s BatchStartStreamsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchStartStreamsRequest) GoString() string {
	return s.String()
}

func (s *BatchStartStreamsRequest) GetId() *string {
	return s.Id
}

func (s *BatchStartStreamsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchStartStreamsRequest) SetId(v string) *BatchStartStreamsRequest {
	s.Id = &v
	return s
}

func (s *BatchStartStreamsRequest) SetOwnerId(v int64) *BatchStartStreamsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchStartStreamsRequest) Validate() error {
	return dara.Validate(s)
}
