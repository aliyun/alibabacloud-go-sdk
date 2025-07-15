// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeployApiResponseBody
	GetRequestId() *string
}

type DeployApiResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeployApiResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeployApiResponseBody) SetRequestId(v string) *DeployApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeployApiResponseBody) Validate() error {
	return dara.Validate(s)
}
