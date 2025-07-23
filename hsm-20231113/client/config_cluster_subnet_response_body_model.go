// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterSubnetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigClusterSubnetResponseBody
	GetRequestId() *string
}

type ConfigClusterSubnetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterSubnetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterSubnetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigClusterSubnetResponseBody) SetRequestId(v string) *ConfigClusterSubnetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigClusterSubnetResponseBody) Validate() error {
	return dara.Validate(s)
}
