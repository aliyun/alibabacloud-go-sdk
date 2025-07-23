// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigClusterWhitelistResponseBody
	GetRequestId() *string
}

type ConfigClusterWhitelistResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigClusterWhitelistResponseBody) SetRequestId(v string) *ConfigClusterWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigClusterWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
