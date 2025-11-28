// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelNames(v *ListSupportModelsResponseBodyModelNames) *ListSupportModelsResponseBody
	GetModelNames() *ListSupportModelsResponseBodyModelNames
	SetRequestId(v string) *ListSupportModelsResponseBody
	GetRequestId() *string
}

type ListSupportModelsResponseBody struct {
	// The list of supported model names.
	ModelNames *ListSupportModelsResponseBodyModelNames `json:"ModelNames,omitempty" xml:"ModelNames,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSupportModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupportModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupportModelsResponseBody) GetModelNames() *ListSupportModelsResponseBodyModelNames {
	return s.ModelNames
}

func (s *ListSupportModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupportModelsResponseBody) SetModelNames(v *ListSupportModelsResponseBodyModelNames) *ListSupportModelsResponseBody {
	s.ModelNames = v
	return s
}

func (s *ListSupportModelsResponseBody) SetRequestId(v string) *ListSupportModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupportModelsResponseBody) Validate() error {
	if s.ModelNames != nil {
		if err := s.ModelNames.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSupportModelsResponseBodyModelNames struct {
	ModelNames []*string `json:"modelNames,omitempty" xml:"modelNames,omitempty" type:"Repeated"`
}

func (s ListSupportModelsResponseBodyModelNames) String() string {
	return dara.Prettify(s)
}

func (s ListSupportModelsResponseBodyModelNames) GoString() string {
	return s.String()
}

func (s *ListSupportModelsResponseBodyModelNames) GetModelNames() []*string {
	return s.ModelNames
}

func (s *ListSupportModelsResponseBodyModelNames) SetModelNames(v []*string) *ListSupportModelsResponseBodyModelNames {
	s.ModelNames = v
	return s
}

func (s *ListSupportModelsResponseBodyModelNames) Validate() error {
	return dara.Validate(s)
}
