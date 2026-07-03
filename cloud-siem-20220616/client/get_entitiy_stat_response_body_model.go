// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEntitiyStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetEntitiyStatResponseBody
	GetCode() *int32
	SetData(v []*GetEntitiyStatResponseBodyData) *GetEntitiyStatResponseBody
	GetData() []*GetEntitiyStatResponseBodyData
	SetMessage(v string) *GetEntitiyStatResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEntitiyStatResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEntitiyStatResponseBody
	GetSuccess() *bool
}

type GetEntitiyStatResponseBody struct {
	// The status code of the request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// 123456
	Data []*GetEntitiyStatResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: successful.
	//
	// - false: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEntitiyStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEntitiyStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetEntitiyStatResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetEntitiyStatResponseBody) GetData() []*GetEntitiyStatResponseBodyData {
	return s.Data
}

func (s *GetEntitiyStatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEntitiyStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEntitiyStatResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEntitiyStatResponseBody) SetCode(v int32) *GetEntitiyStatResponseBody {
	s.Code = &v
	return s
}

func (s *GetEntitiyStatResponseBody) SetData(v []*GetEntitiyStatResponseBodyData) *GetEntitiyStatResponseBody {
	s.Data = v
	return s
}

func (s *GetEntitiyStatResponseBody) SetMessage(v string) *GetEntitiyStatResponseBody {
	s.Message = &v
	return s
}

func (s *GetEntitiyStatResponseBody) SetRequestId(v string) *GetEntitiyStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEntitiyStatResponseBody) SetSuccess(v bool) *GetEntitiyStatResponseBody {
	s.Success = &v
	return s
}

func (s *GetEntitiyStatResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEntitiyStatResponseBodyData struct {
	// The number of entities.
	//
	// example:
	//
	// 3
	EntityNum *int32 `json:"EntityNum,omitempty" xml:"EntityNum,omitempty"`
	// The entity type.
	//
	// example:
	//
	// ip
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity UUID.
	//
	// example:
	//
	// 5cde2118666ffda40783ebd7cec9a60a
	EntityUuid *string `json:"EntityUuid,omitempty" xml:"EntityUuid,omitempty"`
}

func (s GetEntitiyStatResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEntitiyStatResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEntitiyStatResponseBodyData) GetEntityNum() *int32 {
	return s.EntityNum
}

func (s *GetEntitiyStatResponseBodyData) GetEntityType() *string {
	return s.EntityType
}

func (s *GetEntitiyStatResponseBodyData) GetEntityUuid() *string {
	return s.EntityUuid
}

func (s *GetEntitiyStatResponseBodyData) SetEntityNum(v int32) *GetEntitiyStatResponseBodyData {
	s.EntityNum = &v
	return s
}

func (s *GetEntitiyStatResponseBodyData) SetEntityType(v string) *GetEntitiyStatResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *GetEntitiyStatResponseBodyData) SetEntityUuid(v string) *GetEntitiyStatResponseBodyData {
	s.EntityUuid = &v
	return s
}

func (s *GetEntitiyStatResponseBodyData) Validate() error {
	return dara.Validate(s)
}
