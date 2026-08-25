// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateImageResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateImageResponseBody
	GetSuccess() *bool
}

type UpdateImageResponseBody struct {
	// The result of the API request.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateImageResponseBody) SetData(v bool) *UpdateImageResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateImageResponseBody) SetRequestId(v string) *UpdateImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateImageResponseBody) SetSuccess(v bool) *UpdateImageResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateImageResponseBody) Validate() error {
	return dara.Validate(s)
}
