// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusUserSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePrometheusUserSettingResponseBody
	GetRequestId() *string
}

type UpdatePrometheusUserSettingResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePrometheusUserSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusUserSettingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusUserSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusUserSettingResponseBody) SetRequestId(v string) *UpdatePrometheusUserSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusUserSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
