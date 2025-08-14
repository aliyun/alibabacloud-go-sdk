// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveEditingIndexFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIndexFile(v string) *GetLiveEditingIndexFileResponseBody
	GetIndexFile() *string
	SetRequestId(v string) *GetLiveEditingIndexFileResponseBody
	GetRequestId() *string
}

type GetLiveEditingIndexFileResponseBody struct {
	// The URL of the index file.
	IndexFile *string `json:"IndexFile,omitempty" xml:"IndexFile,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E84BE44-58A7-****-****-FBEBEA16EF94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLiveEditingIndexFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLiveEditingIndexFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveEditingIndexFileResponseBody) GetIndexFile() *string {
	return s.IndexFile
}

func (s *GetLiveEditingIndexFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLiveEditingIndexFileResponseBody) SetIndexFile(v string) *GetLiveEditingIndexFileResponseBody {
	s.IndexFile = &v
	return s
}

func (s *GetLiveEditingIndexFileResponseBody) SetRequestId(v string) *GetLiveEditingIndexFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLiveEditingIndexFileResponseBody) Validate() error {
	return dara.Validate(s)
}
