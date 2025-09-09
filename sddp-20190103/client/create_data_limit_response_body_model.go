// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int32) *CreateDataLimitResponseBody
	GetId() *int32
	SetRequestId(v string) *CreateDataLimitResponseBody
	GetRequestId() *string
}

type CreateDataLimitResponseBody struct {
	// The ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLimitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataLimitResponseBody) GetId() *int32 {
	return s.Id
}

func (s *CreateDataLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataLimitResponseBody) SetId(v int32) *CreateDataLimitResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataLimitResponseBody) SetRequestId(v string) *CreateDataLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
