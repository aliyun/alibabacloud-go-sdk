// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigInstanceRemarkResponseBody
	GetRequestId() *string
}

type ConfigInstanceRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigInstanceRemarkResponseBody) SetRequestId(v string) *ConfigInstanceRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
