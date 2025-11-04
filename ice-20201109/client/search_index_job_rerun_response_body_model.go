// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchIndexJobRerunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchIndexJobRerunResponseBody
	GetCode() *string
	SetData(v *SearchIndexJobRerunResponseBodyData) *SearchIndexJobRerunResponseBody
	GetData() *SearchIndexJobRerunResponseBodyData
	SetRequestId(v string) *SearchIndexJobRerunResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SearchIndexJobRerunResponseBody
	GetSuccess() *string
}

type SearchIndexJobRerunResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *SearchIndexJobRerunResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SearchIndexJobRerunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchIndexJobRerunResponseBody) GoString() string {
	return s.String()
}

func (s *SearchIndexJobRerunResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchIndexJobRerunResponseBody) GetData() *SearchIndexJobRerunResponseBodyData {
	return s.Data
}

func (s *SearchIndexJobRerunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchIndexJobRerunResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SearchIndexJobRerunResponseBody) SetCode(v string) *SearchIndexJobRerunResponseBody {
	s.Code = &v
	return s
}

func (s *SearchIndexJobRerunResponseBody) SetData(v *SearchIndexJobRerunResponseBodyData) *SearchIndexJobRerunResponseBody {
	s.Data = v
	return s
}

func (s *SearchIndexJobRerunResponseBody) SetRequestId(v string) *SearchIndexJobRerunResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchIndexJobRerunResponseBody) SetSuccess(v string) *SearchIndexJobRerunResponseBody {
	s.Success = &v
	return s
}

func (s *SearchIndexJobRerunResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchIndexJobRerunResponseBodyData struct {
	// The media asset IDs that do not exist.
	MediaIdsNoExist []*string `json:"MediaIdsNoExist,omitempty" xml:"MediaIdsNoExist,omitempty" type:"Repeated"`
}

func (s SearchIndexJobRerunResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchIndexJobRerunResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchIndexJobRerunResponseBodyData) GetMediaIdsNoExist() []*string {
	return s.MediaIdsNoExist
}

func (s *SearchIndexJobRerunResponseBodyData) SetMediaIdsNoExist(v []*string) *SearchIndexJobRerunResponseBodyData {
	s.MediaIdsNoExist = v
	return s
}

func (s *SearchIndexJobRerunResponseBodyData) Validate() error {
	return dara.Validate(s)
}
