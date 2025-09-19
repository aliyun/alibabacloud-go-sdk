// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSasInstallCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SasInstallCodeResponseBody
	GetRequestId() *string
	SetData(v string) *SasInstallCodeResponseBody
	GetData() *string
}

type SasInstallCodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B256A525-7E42-4BB9-A27C-9017FDDFF1A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The installation verification code that is used to run the installation command when you manually install the Security Center agent.
	//
	// example:
	//
	// eD****
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s SasInstallCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SasInstallCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SasInstallCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SasInstallCodeResponseBody) GetData() *string {
	return s.Data
}

func (s *SasInstallCodeResponseBody) SetRequestId(v string) *SasInstallCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SasInstallCodeResponseBody) SetData(v string) *SasInstallCodeResponseBody {
	s.Data = &v
	return s
}

func (s *SasInstallCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
