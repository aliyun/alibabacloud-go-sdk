// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceRCInstanceSystemDiskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReplaceRCInstanceSystemDiskResponseBody
	GetRequestId() *string
}

type ReplaceRCInstanceSystemDiskResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8B993DA9-5272-5414-94E3-4CA8BA0146C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceRCInstanceSystemDiskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceRCInstanceSystemDiskResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceRCInstanceSystemDiskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceRCInstanceSystemDiskResponseBody) SetRequestId(v string) *ReplaceRCInstanceSystemDiskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceRCInstanceSystemDiskResponseBody) Validate() error {
	return dara.Validate(s)
}
