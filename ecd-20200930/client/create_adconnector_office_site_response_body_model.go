// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateADConnectorOfficeSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *CreateADConnectorOfficeSiteResponseBody
	GetOfficeSiteId() *string
	SetRequestId(v string) *CreateADConnectorOfficeSiteResponseBody
	GetRequestId() *string
}

type CreateADConnectorOfficeSiteResponseBody struct {
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateADConnectorOfficeSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateADConnectorOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateADConnectorOfficeSiteResponseBody) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateADConnectorOfficeSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateADConnectorOfficeSiteResponseBody) SetOfficeSiteId(v string) *CreateADConnectorOfficeSiteResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteResponseBody) SetRequestId(v string) *CreateADConnectorOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateADConnectorOfficeSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
