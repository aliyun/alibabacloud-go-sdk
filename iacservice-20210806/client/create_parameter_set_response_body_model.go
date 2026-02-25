// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateParameterSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetParameterSetId(v string) *CreateParameterSetResponseBody
	GetParameterSetId() *string
	SetRequestId(v string) *CreateParameterSetResponseBody
	GetRequestId() *string
}

type CreateParameterSetResponseBody struct {
	// example:
	//
	// pts-3b6cb9fa4751afff89a4b73779e0d
	ParameterSetId *string `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	// example:
	//
	// 7FA0FF4A-ABD4-54F6-BEAC-B4273EBA10A2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateParameterSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateParameterSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterSetResponseBody) GetParameterSetId() *string {
	return s.ParameterSetId
}

func (s *CreateParameterSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateParameterSetResponseBody) SetParameterSetId(v string) *CreateParameterSetResponseBody {
	s.ParameterSetId = &v
	return s
}

func (s *CreateParameterSetResponseBody) SetRequestId(v string) *CreateParameterSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateParameterSetResponseBody) Validate() error {
	return dara.Validate(s)
}
