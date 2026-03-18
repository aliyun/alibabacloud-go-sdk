// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryApplicationAccessIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationAccessId(v string) *QueryApplicationAccessIdRequest
	GetApplicationAccessId() *string
}

type QueryApplicationAccessIdRequest struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
}

func (s QueryApplicationAccessIdRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryApplicationAccessIdRequest) GoString() string {
	return s.String()
}

func (s *QueryApplicationAccessIdRequest) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *QueryApplicationAccessIdRequest) SetApplicationAccessId(v string) *QueryApplicationAccessIdRequest {
	s.ApplicationAccessId = &v
	return s
}

func (s *QueryApplicationAccessIdRequest) Validate() error {
	return dara.Validate(s)
}
