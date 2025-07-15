// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyADConnectorOfficeSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyADConnectorOfficeSiteResponseBody
	GetRequestId() *string
}

type ModifyADConnectorOfficeSiteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyADConnectorOfficeSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyADConnectorOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyADConnectorOfficeSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyADConnectorOfficeSiteResponseBody) SetRequestId(v string) *ModifyADConnectorOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyADConnectorOfficeSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
