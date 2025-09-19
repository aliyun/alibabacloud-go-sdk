// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVulConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyVulConfigResponseBody
	GetRequestId() *string
}

type ModifyVulConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 24DDBE06-58FF-5E5E-9241-D2010D7913C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVulConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyVulConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVulConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyVulConfigResponseBody) SetRequestId(v string) *ModifyVulConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyVulConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
