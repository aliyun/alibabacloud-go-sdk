// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExplorerModuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExplorerModuleAttributeResponseBody
	GetRequestId() *string
}

type UpdateExplorerModuleAttributeResponseBody struct {
	// example:
	//
	// 25B274BA-E672-58C0-8602-541281B6F758
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateExplorerModuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExplorerModuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExplorerModuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExplorerModuleAttributeResponseBody) SetRequestId(v string) *UpdateExplorerModuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExplorerModuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
