// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DisableImageResponseBody
	GetData() *bool
	SetRequestId(v string) *DisableImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableImageResponseBody
	GetSuccess() *bool
}

type DisableImageResponseBody struct {
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

func (s DisableImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableImageResponseBody) GoString() string {
	return s.String()
}

func (s *DisableImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DisableImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableImageResponseBody) SetData(v bool) *DisableImageResponseBody {
	s.Data = &v
	return s
}

func (s *DisableImageResponseBody) SetRequestId(v string) *DisableImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableImageResponseBody) SetSuccess(v bool) *DisableImageResponseBody {
	s.Success = &v
	return s
}

func (s *DisableImageResponseBody) Validate() error {
	return dara.Validate(s)
}
