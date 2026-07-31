// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenameSemanticViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenameSemanticViewResponseBody
	GetRequestId() *string
}

type RenameSemanticViewResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenameSemanticViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenameSemanticViewResponseBody) GoString() string {
	return s.String()
}

func (s *RenameSemanticViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenameSemanticViewResponseBody) SetRequestId(v string) *RenameSemanticViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameSemanticViewResponseBody) Validate() error {
	return dara.Validate(s)
}
