// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigImageRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigImageRemarkResponseBody
	GetRequestId() *string
}

type ConfigImageRemarkResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigImageRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigImageRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigImageRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigImageRemarkResponseBody) SetRequestId(v string) *ConfigImageRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigImageRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
