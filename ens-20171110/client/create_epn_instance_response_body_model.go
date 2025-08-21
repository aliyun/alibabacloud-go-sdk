// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEpnInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEPNInstanceId(v string) *CreateEpnInstanceResponseBody
	GetEPNInstanceId() *string
	SetRequestId(v string) *CreateEpnInstanceResponseBody
	GetRequestId() *string
}

type CreateEpnInstanceResponseBody struct {
	// The ID of the EPN instance.
	//
	// example:
	//
	// epn-xxxxx
	EPNInstanceId *string `json:"EPNInstanceId,omitempty" xml:"EPNInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEpnInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEpnInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEpnInstanceResponseBody) GetEPNInstanceId() *string {
	return s.EPNInstanceId
}

func (s *CreateEpnInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEpnInstanceResponseBody) SetEPNInstanceId(v string) *CreateEpnInstanceResponseBody {
	s.EPNInstanceId = &v
	return s
}

func (s *CreateEpnInstanceResponseBody) SetRequestId(v string) *CreateEpnInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEpnInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
