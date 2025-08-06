// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIStatisTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAIStatisTypeResponseBody
	GetRequestId() *string
	SetTypes(v string) *ListAIStatisTypeResponseBody
	GetTypes() *string
}

type ListAIStatisTypeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Types     *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListAIStatisTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIStatisTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIStatisTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIStatisTypeResponseBody) GetTypes() *string {
	return s.Types
}

func (s *ListAIStatisTypeResponseBody) SetRequestId(v string) *ListAIStatisTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIStatisTypeResponseBody) SetTypes(v string) *ListAIStatisTypeResponseBody {
	s.Types = &v
	return s
}

func (s *ListAIStatisTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
