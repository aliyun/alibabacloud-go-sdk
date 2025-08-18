// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpsBasicConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateHttpsBasicConfigurationResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateHttpsBasicConfigurationResponseBody
	GetRequestId() *string
}

type CreateHttpsBasicConfigurationResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHttpsBasicConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpsBasicConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpsBasicConfigurationResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateHttpsBasicConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpsBasicConfigurationResponseBody) SetConfigId(v int64) *CreateHttpsBasicConfigurationResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateHttpsBasicConfigurationResponseBody) SetRequestId(v string) *CreateHttpsBasicConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpsBasicConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
