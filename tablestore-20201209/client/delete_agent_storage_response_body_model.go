// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgentStorageResponseBody
	GetRequestId() *string
}

type DeleteAgentStorageResponseBody struct {
	// The request ID, which can be used to troubleshoot and locate issues.
	//
	// example:
	//
	// E734979F-5A44-5993-9CE5-C23103576923
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgentStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentStorageResponseBody) SetRequestId(v string) *DeleteAgentStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
