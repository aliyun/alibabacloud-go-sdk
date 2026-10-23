// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKibanaSettingsResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateKibanaSettingsResponseBody
	GetResult() *bool
}

type UpdateKibanaSettingsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DC*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result. Valid values:
	//
	// - true: The Kibana language was modified.
	//
	// - false: The Kibana language failed to be modified.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateKibanaSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKibanaSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKibanaSettingsResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateKibanaSettingsResponseBody) SetRequestId(v string) *UpdateKibanaSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKibanaSettingsResponseBody) SetResult(v bool) *UpdateKibanaSettingsResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateKibanaSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
