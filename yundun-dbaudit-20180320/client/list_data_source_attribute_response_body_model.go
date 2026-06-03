// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbList(v []*ListDataSourceAttributeResponseBodyDbList) *ListDataSourceAttributeResponseBody
	GetDbList() []*ListDataSourceAttributeResponseBodyDbList
	SetRequestId(v string) *ListDataSourceAttributeResponseBody
	GetRequestId() *string
}

type ListDataSourceAttributeResponseBody struct {
	DbList []*ListDataSourceAttributeResponseBodyDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	// example:
	//
	// 1B217656-2267-4FBF-B26C-CDD201BDC3B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeResponseBody) GetDbList() []*ListDataSourceAttributeResponseBodyDbList {
	return s.DbList
}

func (s *ListDataSourceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceAttributeResponseBody) SetDbList(v []*ListDataSourceAttributeResponseBodyDbList) *ListDataSourceAttributeResponseBody {
	s.DbList = v
	return s
}

func (s *ListDataSourceAttributeResponseBody) SetRequestId(v string) *ListDataSourceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceAttributeResponseBody) Validate() error {
	if s.DbList != nil {
		for _, item := range s.DbList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourceAttributeResponseBodyDbList struct {
	// example:
	//
	// All
	AuditMode *string `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	// example:
	//
	// 1
	DbId          *int32  `json:"DbId,omitempty" xml:"DbId,omitempty"`
	FreqAuditMode *string `json:"FreqAuditMode,omitempty" xml:"FreqAuditMode,omitempty"`
	// example:
	//
	// 100
	ResultAuditMaxLine *int32 `json:"ResultAuditMaxLine,omitempty" xml:"ResultAuditMaxLine,omitempty"`
	// example:
	//
	// 10
	ResultAuditMaxSize *int32 `json:"ResultAuditMaxSize,omitempty" xml:"ResultAuditMaxSize,omitempty"`
	// example:
	//
	// Close
	ResultAuditMode *string `json:"ResultAuditMode,omitempty" xml:"ResultAuditMode,omitempty"`
}

func (s ListDataSourceAttributeResponseBodyDbList) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceAttributeResponseBodyDbList) GoString() string {
	return s.String()
}

func (s *ListDataSourceAttributeResponseBodyDbList) GetAuditMode() *string {
	return s.AuditMode
}

func (s *ListDataSourceAttributeResponseBodyDbList) GetDbId() *int32 {
	return s.DbId
}

func (s *ListDataSourceAttributeResponseBodyDbList) GetFreqAuditMode() *string {
	return s.FreqAuditMode
}

func (s *ListDataSourceAttributeResponseBodyDbList) GetResultAuditMaxLine() *int32 {
	return s.ResultAuditMaxLine
}

func (s *ListDataSourceAttributeResponseBodyDbList) GetResultAuditMaxSize() *int32 {
	return s.ResultAuditMaxSize
}

func (s *ListDataSourceAttributeResponseBodyDbList) GetResultAuditMode() *string {
	return s.ResultAuditMode
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetAuditMode(v string) *ListDataSourceAttributeResponseBodyDbList {
	s.AuditMode = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetDbId(v int32) *ListDataSourceAttributeResponseBodyDbList {
	s.DbId = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetFreqAuditMode(v string) *ListDataSourceAttributeResponseBodyDbList {
	s.FreqAuditMode = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetResultAuditMaxLine(v int32) *ListDataSourceAttributeResponseBodyDbList {
	s.ResultAuditMaxLine = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetResultAuditMaxSize(v int32) *ListDataSourceAttributeResponseBodyDbList {
	s.ResultAuditMaxSize = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) SetResultAuditMode(v string) *ListDataSourceAttributeResponseBodyDbList {
	s.ResultAuditMode = &v
	return s
}

func (s *ListDataSourceAttributeResponseBodyDbList) Validate() error {
	return dara.Validate(s)
}
