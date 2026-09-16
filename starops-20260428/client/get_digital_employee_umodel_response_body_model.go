// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeUmodelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommonSchemaRef(v []*GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) *GetDigitalEmployeeUmodelResponseBody
	GetCommonSchemaRef() []*GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef
	SetDescription(v string) *GetDigitalEmployeeUmodelResponseBody
	GetDescription() *string
	SetName(v string) *GetDigitalEmployeeUmodelResponseBody
	GetName() *string
	SetRequestId(v string) *GetDigitalEmployeeUmodelResponseBody
	GetRequestId() *string
}

type GetDigitalEmployeeUmodelResponseBody struct {
	// The common schemas referenced by the digital employee UModel.
	CommonSchemaRef []*GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef `json:"commonSchemaRef,omitempty" xml:"commonSchemaRef,omitempty" type:"Repeated"`
	// The description of the digital employee UModel.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the digital employee to which the UModel belongs.
	//
	// example:
	//
	// sample-agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-1234567890AB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDigitalEmployeeUmodelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeUmodelResponseBody) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeUmodelResponseBody) GetCommonSchemaRef() []*GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef {
	return s.CommonSchemaRef
}

func (s *GetDigitalEmployeeUmodelResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDigitalEmployeeUmodelResponseBody) GetName() *string {
	return s.Name
}

func (s *GetDigitalEmployeeUmodelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDigitalEmployeeUmodelResponseBody) SetCommonSchemaRef(v []*GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) *GetDigitalEmployeeUmodelResponseBody {
	s.CommonSchemaRef = v
	return s
}

func (s *GetDigitalEmployeeUmodelResponseBody) SetDescription(v string) *GetDigitalEmployeeUmodelResponseBody {
	s.Description = &v
	return s
}

func (s *GetDigitalEmployeeUmodelResponseBody) SetName(v string) *GetDigitalEmployeeUmodelResponseBody {
	s.Name = &v
	return s
}

func (s *GetDigitalEmployeeUmodelResponseBody) SetRequestId(v string) *GetDigitalEmployeeUmodelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDigitalEmployeeUmodelResponseBody) Validate() error {
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

type GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef struct {
	// The schema group.
	//
	// example:
	//
	// default
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The schema version.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) GetGroup() *string {
	return s.Group
}

func (s *GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) GetVersion() *string {
	return s.Version
}

func (s *GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) SetGroup(v string) *GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef {
	s.Group = &v
	return s
}

func (s *GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) SetVersion(v string) *GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef {
	s.Version = &v
	return s
}

func (s *GetDigitalEmployeeUmodelResponseBodyCommonSchemaRef) Validate() error {
	return dara.Validate(s)
}
