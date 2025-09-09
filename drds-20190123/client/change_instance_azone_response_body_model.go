// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeInstanceAzoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeInstanceAzoneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeInstanceAzoneResponseBody
	GetSuccess() *bool
}

type ChangeInstanceAzoneResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4305162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeInstanceAzoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeInstanceAzoneResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeInstanceAzoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeInstanceAzoneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeInstanceAzoneResponseBody) SetRequestId(v string) *ChangeInstanceAzoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeInstanceAzoneResponseBody) SetSuccess(v bool) *ChangeInstanceAzoneResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeInstanceAzoneResponseBody) Validate() error {
	return dara.Validate(s)
}
