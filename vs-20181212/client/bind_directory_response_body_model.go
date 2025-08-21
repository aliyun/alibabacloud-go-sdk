// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BindDirectoryResponseBody
	GetRequestId() *string
}

type BindDirectoryResponseBody struct {
	// example:
	//
	// 3CB843A9-DD34-4881-B8D6-B0D539D111E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *BindDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindDirectoryResponseBody) SetRequestId(v string) *BindDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
