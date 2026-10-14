// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAifgUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAifgUrlResponseBody
	GetCode() *string
	SetData(v *GetAifgUrlResponseBodyData) *GetAifgUrlResponseBody
	GetData() *GetAifgUrlResponseBodyData
	SetMessage(v string) *GetAifgUrlResponseBody
	GetMessage() *string
}

type GetAifgUrlResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetAifgUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The description of the status code.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAifgUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAifgUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAifgUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAifgUrlResponseBody) GetData() *GetAifgUrlResponseBodyData {
	return s.Data
}

func (s *GetAifgUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAifgUrlResponseBody) SetCode(v string) *GetAifgUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAifgUrlResponseBody) SetData(v *GetAifgUrlResponseBodyData) *GetAifgUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAifgUrlResponseBody) SetMessage(v string) *GetAifgUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAifgUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAifgUrlResponseBodyData struct {
	// The URL of the diagnostic details.
	//
	// example:
	//
	// /diagnose/detail/qe3Z34sa
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAifgUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAifgUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAifgUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetAifgUrlResponseBodyData) SetUrl(v string) *GetAifgUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetAifgUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
