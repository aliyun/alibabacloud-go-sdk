// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iControlVodAppServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppServiceStatus(v string) *ControlVodAppServiceResponseBody
	GetAppServiceStatus() *string
	SetRequestId(v string) *ControlVodAppServiceResponseBody
	GetRequestId() *string
}

type ControlVodAppServiceResponseBody struct {
	AppServiceStatus *string `json:"AppServiceStatus,omitempty" xml:"AppServiceStatus,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ControlVodAppServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ControlVodAppServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ControlVodAppServiceResponseBody) GetAppServiceStatus() *string {
	return s.AppServiceStatus
}

func (s *ControlVodAppServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ControlVodAppServiceResponseBody) SetAppServiceStatus(v string) *ControlVodAppServiceResponseBody {
	s.AppServiceStatus = &v
	return s
}

func (s *ControlVodAppServiceResponseBody) SetRequestId(v string) *ControlVodAppServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ControlVodAppServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
