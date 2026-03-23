// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSendifyTrialServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoLoginURL(v string) *OpenSendifyTrialServiceResponseBody
	GetAutoLoginURL() *string
	SetRequestId(v string) *OpenSendifyTrialServiceResponseBody
	GetRequestId() *string
}

type OpenSendifyTrialServiceResponseBody struct {
	// example:
	//
	// http
	AutoLoginURL *string `json:"AutoLoginURL,omitempty" xml:"AutoLoginURL,omitempty"`
	// example:
	//
	// 1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenSendifyTrialServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenSendifyTrialServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenSendifyTrialServiceResponseBody) GetAutoLoginURL() *string {
	return s.AutoLoginURL
}

func (s *OpenSendifyTrialServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenSendifyTrialServiceResponseBody) SetAutoLoginURL(v string) *OpenSendifyTrialServiceResponseBody {
	s.AutoLoginURL = &v
	return s
}

func (s *OpenSendifyTrialServiceResponseBody) SetRequestId(v string) *OpenSendifyTrialServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenSendifyTrialServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
