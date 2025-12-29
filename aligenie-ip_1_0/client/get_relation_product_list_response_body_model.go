// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRelationProductListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *GetRelationProductListResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *GetRelationProductListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRelationProductListResponseBody
	GetRequestId() *string
	SetResult(v []*GetRelationProductListResponseBodyResult) *GetRelationProductListResponseBody
	GetResult() []*GetRelationProductListResponseBodyResult
	SetStatusCode(v int32) *GetRelationProductListResponseBody
	GetStatusCode() *int32
}

type GetRelationProductListResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetRelationProductListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetRelationProductListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRelationProductListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRelationProductListResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *GetRelationProductListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRelationProductListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRelationProductListResponseBody) GetResult() []*GetRelationProductListResponseBodyResult {
	return s.Result
}

func (s *GetRelationProductListResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRelationProductListResponseBody) SetExtentions(v map[string]interface{}) *GetRelationProductListResponseBody {
	s.Extentions = v
	return s
}

func (s *GetRelationProductListResponseBody) SetMessage(v string) *GetRelationProductListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRelationProductListResponseBody) SetRequestId(v string) *GetRelationProductListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRelationProductListResponseBody) SetResult(v []*GetRelationProductListResponseBodyResult) *GetRelationProductListResponseBody {
	s.Result = v
	return s
}

func (s *GetRelationProductListResponseBody) SetStatusCode(v int32) *GetRelationProductListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetRelationProductListResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRelationProductListResponseBodyResult struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// jTOSl***l1odxImRw
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetRelationProductListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetRelationProductListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRelationProductListResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetRelationProductListResponseBodyResult) GetPk() *string {
	return s.Pk
}

func (s *GetRelationProductListResponseBodyResult) SetName(v string) *GetRelationProductListResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetRelationProductListResponseBodyResult) SetPk(v string) *GetRelationProductListResponseBodyResult {
	s.Pk = &v
	return s
}

func (s *GetRelationProductListResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
