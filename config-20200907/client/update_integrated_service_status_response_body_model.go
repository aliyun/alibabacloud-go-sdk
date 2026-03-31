// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegratedServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIntegratedServiceStatusResponseBody
	GetRequestId() *string
}

type UpdateIntegratedServiceStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 02EB7638-B029-5ABB-93F5-A2ABEEAC282D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIntegratedServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegratedServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntegratedServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIntegratedServiceStatusResponseBody) SetRequestId(v string) *UpdateIntegratedServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIntegratedServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
