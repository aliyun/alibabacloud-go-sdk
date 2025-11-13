// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleSimilarMaliciousFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *HandleSimilarMaliciousFilesResponseBody
	GetData() *int64
	SetRequestId(v string) *HandleSimilarMaliciousFilesResponseBody
	GetRequestId() *string
}

type HandleSimilarMaliciousFilesResponseBody struct {
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 8C376***AE74FB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s HandleSimilarMaliciousFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleSimilarMaliciousFilesResponseBody) GoString() string {
	return s.String()
}

func (s *HandleSimilarMaliciousFilesResponseBody) GetData() *int64 {
	return s.Data
}

func (s *HandleSimilarMaliciousFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleSimilarMaliciousFilesResponseBody) SetData(v int64) *HandleSimilarMaliciousFilesResponseBody {
	s.Data = &v
	return s
}

func (s *HandleSimilarMaliciousFilesResponseBody) SetRequestId(v string) *HandleSimilarMaliciousFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleSimilarMaliciousFilesResponseBody) Validate() error {
	return dara.Validate(s)
}
