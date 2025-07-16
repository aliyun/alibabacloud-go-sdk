// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSheetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateSheetResponseBody
	GetId() *string
	SetName(v string) *CreateSheetResponseBody
	GetName() *string
	SetRequestId(v string) *CreateSheetResponseBody
	GetRequestId() *string
	SetVisibility(v string) *CreateSheetResponseBody
	GetVisibility() *string
}

type CreateSheetResponseBody struct {
	// example:
	//
	// stxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// Sheet1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// visible
	Visibility *string `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (s CreateSheetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSheetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSheetResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateSheetResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateSheetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSheetResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *CreateSheetResponseBody) SetId(v string) *CreateSheetResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSheetResponseBody) SetName(v string) *CreateSheetResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSheetResponseBody) SetRequestId(v string) *CreateSheetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSheetResponseBody) SetVisibility(v string) *CreateSheetResponseBody {
	s.Visibility = &v
	return s
}

func (s *CreateSheetResponseBody) Validate() error {
	return dara.Validate(s)
}
