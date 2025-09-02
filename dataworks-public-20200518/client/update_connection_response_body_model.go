// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateConnectionResponseBody
	GetData() *bool
	SetHttpStatusCode(v string) *UpdateConnectionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateConnectionResponseBody
	GetSuccess() *bool
}

type UpdateConnectionResponseBody struct {
	// Indicates whether the data source is updated.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 0bc14115159376359****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConnectionResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateConnectionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateConnectionResponseBody) SetData(v bool) *UpdateConnectionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetHttpStatusCode(v string) *UpdateConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetRequestId(v string) *UpdateConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateConnectionResponseBody) SetSuccess(v bool) *UpdateConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
