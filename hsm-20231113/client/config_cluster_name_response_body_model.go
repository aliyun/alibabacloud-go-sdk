// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigClusterNameResponseBody
	GetRequestId() *string
}

type ConfigClusterNameResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterNameResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigClusterNameResponseBody) SetRequestId(v string) *ConfigClusterNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigClusterNameResponseBody) Validate() error {
	return dara.Validate(s)
}
