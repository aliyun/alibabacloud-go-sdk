// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCACertificateLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogList(v []*ListCACertificateLogResponseBodyLogList) *ListCACertificateLogResponseBody
	GetLogList() []*ListCACertificateLogResponseBodyLogList
	SetRequestId(v string) *ListCACertificateLogResponseBody
	GetRequestId() *string
}

type ListCACertificateLogResponseBody struct {
	LogList []*ListCACertificateLogResponseBodyLogList `json:"LogList,omitempty" xml:"LogList,omitempty" type:"Repeated"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCACertificateLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCACertificateLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogResponseBody) GetLogList() []*ListCACertificateLogResponseBodyLogList {
	return s.LogList
}

func (s *ListCACertificateLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCACertificateLogResponseBody) SetLogList(v []*ListCACertificateLogResponseBodyLogList) *ListCACertificateLogResponseBody {
	s.LogList = v
	return s
}

func (s *ListCACertificateLogResponseBody) SetRequestId(v string) *ListCACertificateLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCACertificateLogResponseBody) Validate() error {
	if s.LogList != nil {
		for _, item := range s.LogList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCACertificateLogResponseBodyLogList struct {
	// example:
	//
	// add sub-root ca
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1634539509000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// ADD
	OpType *string `json:"OpType,omitempty" xml:"OpType,omitempty"`
}

func (s ListCACertificateLogResponseBodyLogList) String() string {
	return dara.Prettify(s)
}

func (s ListCACertificateLogResponseBodyLogList) GoString() string {
	return s.String()
}

func (s *ListCACertificateLogResponseBodyLogList) GetContent() *string {
	return s.Content
}

func (s *ListCACertificateLogResponseBodyLogList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCACertificateLogResponseBodyLogList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListCACertificateLogResponseBodyLogList) GetOpType() *string {
	return s.OpType
}

func (s *ListCACertificateLogResponseBodyLogList) SetContent(v string) *ListCACertificateLogResponseBodyLogList {
	s.Content = &v
	return s
}

func (s *ListCACertificateLogResponseBodyLogList) SetCreateTime(v int64) *ListCACertificateLogResponseBodyLogList {
	s.CreateTime = &v
	return s
}

func (s *ListCACertificateLogResponseBodyLogList) SetIdentifier(v string) *ListCACertificateLogResponseBodyLogList {
	s.Identifier = &v
	return s
}

func (s *ListCACertificateLogResponseBodyLogList) SetOpType(v string) *ListCACertificateLogResponseBodyLogList {
	s.OpType = &v
	return s
}

func (s *ListCACertificateLogResponseBodyLogList) Validate() error {
	return dara.Validate(s)
}
