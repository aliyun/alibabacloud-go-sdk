// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateConnectionResponseBody
	GetData() *int64
	SetHttpStatusCode(v string) *CreateConnectionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateConnectionResponseBody
	GetSuccess() *bool
}

type CreateConnectionResponseBody struct {
	// The data source ID.
	//
	// example:
	//
	// 123
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc141151593763****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnectionResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateConnectionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateConnectionResponseBody) SetData(v int64) *CreateConnectionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateConnectionResponseBody) SetHttpStatusCode(v string) *CreateConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateConnectionResponseBody) SetRequestId(v string) *CreateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConnectionResponseBody) SetSuccess(v bool) *CreateConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
