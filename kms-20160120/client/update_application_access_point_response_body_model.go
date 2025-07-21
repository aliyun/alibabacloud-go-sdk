// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateApplicationAccessPointResponseBody
	GetRequestId() *string
}

type UpdateApplicationAccessPointResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationAccessPointResponseBody) SetRequestId(v string) *UpdateApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
