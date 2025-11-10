// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUmodelCommonSchemaRefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonSchemaRef(v []*GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) *GetUmodelCommonSchemaRefResponseBody
	GetCommonSchemaRef() []*GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef
}

type GetUmodelCommonSchemaRefResponseBody struct {
	CommonSchemaRef []*GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
}

func (s GetUmodelCommonSchemaRefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelCommonSchemaRefResponseBody) GoString() string {
	return s.String()
}

func (s *GetUmodelCommonSchemaRefResponseBody) GetCommonSchemaRef() []*GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef {
	return s.CommonSchemaRef
}

func (s *GetUmodelCommonSchemaRefResponseBody) SetCommonSchemaRef(v []*GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) *GetUmodelCommonSchemaRefResponseBody {
	s.CommonSchemaRef = v
	return s
}

func (s *GetUmodelCommonSchemaRefResponseBody) Validate() error {
	if s.CommonSchemaRef != nil {
		for _, item := range s.CommonSchemaRef {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef struct {
	// example:
	//
	// apm-common
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// example:
	//
	// 0.1.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) String() string {
	return dara.Prettify(s)
}

func (s GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) GetGroup() *string {
	return s.Group
}

func (s *GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) GetVersion() *string {
	return s.Version
}

func (s *GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) SetGroup(v string) *GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) SetVersion(v string) *GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef {
	s.Version = &v
	return s
}

func (s *GetUmodelCommonSchemaRefResponseBodyCommonSchemaRef) Validate() error {
	return dara.Validate(s)
}
