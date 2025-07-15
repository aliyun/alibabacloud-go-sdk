// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentId(v string) *AddCasterComponentResponseBody
	GetComponentId() *string
	SetRequestId(v string) *AddCasterComponentResponseBody
	GetRequestId() *string
}

type AddCasterComponentResponseBody struct {
	// The component ID. The value can be used as the value of a request parameter to query, modify, or delete a production studio.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCasterComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCasterComponentResponseBody) GoString() string {
	return s.String()
}

func (s *AddCasterComponentResponseBody) GetComponentId() *string {
	return s.ComponentId
}

func (s *AddCasterComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCasterComponentResponseBody) SetComponentId(v string) *AddCasterComponentResponseBody {
	s.ComponentId = &v
	return s
}

func (s *AddCasterComponentResponseBody) SetRequestId(v string) *AddCasterComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCasterComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
