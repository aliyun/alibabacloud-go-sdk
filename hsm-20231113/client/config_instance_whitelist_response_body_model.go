// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigInstanceWhitelistResponseBody
	GetRequestId() *string
}

type ConfigInstanceWhitelistResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigInstanceWhitelistResponseBody) SetRequestId(v string) *ConfigInstanceWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
