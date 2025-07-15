// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterComponentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponentId(v string) *ModifyCasterComponentResponseBody
	GetComponentId() *string
	SetRequestId(v string) *ModifyCasterComponentResponseBody
	GetRequestId() *string
}

type ModifyCasterComponentResponseBody struct {
	// The ID of the component. You can use the ID as a request parameter in the API operation that is called to query components in a production studio.
	//
	// example:
	//
	// 05ab713c-676e-49c0-96ce-cc408da1****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCasterComponentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterComponentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCasterComponentResponseBody) GetComponentId() *string {
	return s.ComponentId
}

func (s *ModifyCasterComponentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCasterComponentResponseBody) SetComponentId(v string) *ModifyCasterComponentResponseBody {
	s.ComponentId = &v
	return s
}

func (s *ModifyCasterComponentResponseBody) SetRequestId(v string) *ModifyCasterComponentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCasterComponentResponseBody) Validate() error {
	return dara.Validate(s)
}
