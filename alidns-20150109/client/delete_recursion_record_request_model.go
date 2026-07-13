// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecursionRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRecursionRecordRequest
	GetClientToken() *string
	SetRecordId(v string) *DeleteRecursionRecordRequest
	GetRecordId() *string
}

type DeleteRecursionRecordRequest struct {
	// The client token that ensures the idempotence of the request. The client generates this value. It must be unique across requests. The value can be up to 64 ASCII characters long.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the DNS record. This is the unique identifier for the record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 17432432424
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
}

func (s DeleteRecursionRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecursionRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecursionRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRecursionRecordRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *DeleteRecursionRecordRequest) SetClientToken(v string) *DeleteRecursionRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRecursionRecordRequest) SetRecordId(v string) *DeleteRecursionRecordRequest {
	s.RecordId = &v
	return s
}

func (s *DeleteRecursionRecordRequest) Validate() error {
	return dara.Validate(s)
}
