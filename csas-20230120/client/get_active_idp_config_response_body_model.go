// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetActiveIdpConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetActiveIdpConfigResponseBodyData) *GetActiveIdpConfigResponseBody
	GetData() *GetActiveIdpConfigResponseBodyData
	SetRequestId(v string) *GetActiveIdpConfigResponseBody
	GetRequestId() *string
}

type GetActiveIdpConfigResponseBody struct {
	Data *GetActiveIdpConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetActiveIdpConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetActiveIdpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetActiveIdpConfigResponseBody) GetData() *GetActiveIdpConfigResponseBodyData {
	return s.Data
}

func (s *GetActiveIdpConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetActiveIdpConfigResponseBody) SetData(v *GetActiveIdpConfigResponseBodyData) *GetActiveIdpConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetActiveIdpConfigResponseBody) SetRequestId(v string) *GetActiveIdpConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetActiveIdpConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetActiveIdpConfigResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// idp-cfg001
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DingTalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetActiveIdpConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetActiveIdpConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetActiveIdpConfigResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetActiveIdpConfigResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetActiveIdpConfigResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetActiveIdpConfigResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetActiveIdpConfigResponseBodyData) SetDescription(v string) *GetActiveIdpConfigResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetActiveIdpConfigResponseBodyData) SetId(v string) *GetActiveIdpConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetActiveIdpConfigResponseBodyData) SetName(v string) *GetActiveIdpConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetActiveIdpConfigResponseBodyData) SetType(v string) *GetActiveIdpConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetActiveIdpConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
