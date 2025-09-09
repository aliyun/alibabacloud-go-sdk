// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateConfigResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateConfigResponseBody
	GetRequestId() *string
}

type CreateConfigResponseBody struct {
	// The ID of the common alert configuration.
	//
	// example:
	//
	// 12300
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 208B016D-4CB9-4A85-96A5-0B8ED1EBF271
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConfigResponseBody) SetId(v int64) *CreateConfigResponseBody {
	s.Id = &v
	return s
}

func (s *CreateConfigResponseBody) SetRequestId(v string) *CreateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
