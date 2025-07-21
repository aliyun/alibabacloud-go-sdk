// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteApplicationAccessPointResponseBody
	GetRequestId() *string
}

type DeleteApplicationAccessPointResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// bcfefe15-46f0-44a3-bd96-3d422474b71a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationAccessPointResponseBody) SetRequestId(v string) *DeleteApplicationAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationAccessPointResponseBody) Validate() error {
	return dara.Validate(s)
}
