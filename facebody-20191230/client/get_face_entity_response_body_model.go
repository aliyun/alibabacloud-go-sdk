// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFaceEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFaceEntityResponseBodyData) *GetFaceEntityResponseBody
	GetData() *GetFaceEntityResponseBodyData
	SetRequestId(v string) *GetFaceEntityResponseBody
	GetRequestId() *string
}

type GetFaceEntityResponseBody struct {
	Data *GetFaceEntityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DA7CAFEB-0A37-4496-9CDF-E3DBB6309CB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFaceEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFaceEntityResponseBody) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBody) GetData() *GetFaceEntityResponseBodyData {
	return s.Data
}

func (s *GetFaceEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFaceEntityResponseBody) SetData(v *GetFaceEntityResponseBodyData) *GetFaceEntityResponseBody {
	s.Data = v
	return s
}

func (s *GetFaceEntityResponseBody) SetRequestId(v string) *GetFaceEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFaceEntityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFaceEntityResponseBodyData struct {
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 66
	EntityId *string                               `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Faces    []*GetFaceEntityResponseBodyDataFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	Labels   *string                               `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s GetFaceEntityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFaceEntityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBodyData) GetDbName() *string {
	return s.DbName
}

func (s *GetFaceEntityResponseBodyData) GetEntityId() *string {
	return s.EntityId
}

func (s *GetFaceEntityResponseBodyData) GetFaces() []*GetFaceEntityResponseBodyDataFaces {
	return s.Faces
}

func (s *GetFaceEntityResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *GetFaceEntityResponseBodyData) SetDbName(v string) *GetFaceEntityResponseBodyData {
	s.DbName = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetEntityId(v string) *GetFaceEntityResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetFaces(v []*GetFaceEntityResponseBodyDataFaces) *GetFaceEntityResponseBodyData {
	s.Faces = v
	return s
}

func (s *GetFaceEntityResponseBodyData) SetLabels(v string) *GetFaceEntityResponseBodyData {
	s.Labels = &v
	return s
}

func (s *GetFaceEntityResponseBodyData) Validate() error {
	if s.Faces != nil {
		for _, item := range s.Faces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFaceEntityResponseBodyDataFaces struct {
	// example:
	//
	// 3
	FaceId *string `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
}

func (s GetFaceEntityResponseBodyDataFaces) String() string {
	return dara.Prettify(s)
}

func (s GetFaceEntityResponseBodyDataFaces) GoString() string {
	return s.String()
}

func (s *GetFaceEntityResponseBodyDataFaces) GetFaceId() *string {
	return s.FaceId
}

func (s *GetFaceEntityResponseBodyDataFaces) SetFaceId(v string) *GetFaceEntityResponseBodyDataFaces {
	s.FaceId = &v
	return s
}

func (s *GetFaceEntityResponseBodyDataFaces) Validate() error {
	return dara.Validate(s)
}
