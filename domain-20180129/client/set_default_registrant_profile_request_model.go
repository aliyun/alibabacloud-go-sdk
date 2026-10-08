// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultRegistrantProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistrantProfileId(v int64) *SetDefaultRegistrantProfileRequest
	GetRegistrantProfileId() *int64
	SetUserClientIp(v string) *SetDefaultRegistrantProfileRequest
	GetUserClientIp() *string
}

type SetDefaultRegistrantProfileRequest struct {
	// The ID of the contact template to be set as default.
	//
	// The system automatically generates this ID after the template is successfully created. You can invoke the [QueryRegistrantProfiles](https://help.aliyun.com/document_detail/67701.html) API to query the template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// The user IP address. The default value is **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SetDefaultRegistrantProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultRegistrantProfileRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultRegistrantProfileRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *SetDefaultRegistrantProfileRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SetDefaultRegistrantProfileRequest) SetRegistrantProfileId(v int64) *SetDefaultRegistrantProfileRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *SetDefaultRegistrantProfileRequest) SetUserClientIp(v string) *SetDefaultRegistrantProfileRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDefaultRegistrantProfileRequest) Validate() error {
	return dara.Validate(s)
}
