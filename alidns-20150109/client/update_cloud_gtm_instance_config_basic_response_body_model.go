// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigBasicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudGtmInstanceConfigBasicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCloudGtmInstanceConfigBasicResponseBody
	GetSuccess() *bool
}

type UpdateCloudGtmInstanceConfigBasicResponseBody struct {
	// Unique request identification code.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful:
	//
	// - **true**: The call succeeded.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigBasicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigBasicResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigBasicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudGtmInstanceConfigBasicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCloudGtmInstanceConfigBasicResponseBody) SetRequestId(v string) *UpdateCloudGtmInstanceConfigBasicResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicResponseBody) SetSuccess(v bool) *UpdateCloudGtmInstanceConfigBasicResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicResponseBody) Validate() error {
	return dara.Validate(s)
}
