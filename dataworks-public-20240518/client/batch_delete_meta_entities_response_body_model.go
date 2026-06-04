// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteMetaEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchDeleteMetaEntitiesResponseBody
	GetRequestId() *string
	SetResults(v []*MetaEntityWriteResult) *BatchDeleteMetaEntitiesResponseBody
	GetResults() []*MetaEntityWriteResult
	SetSuccess(v bool) *BatchDeleteMetaEntitiesResponseBody
	GetSuccess() *bool
}

type BatchDeleteMetaEntitiesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1FC02D76-4A94-5D97-B52C-00A031B95359
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*MetaEntityWriteResult `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// []
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteMetaEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteMetaEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteMetaEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteMetaEntitiesResponseBody) GetResults() []*MetaEntityWriteResult {
	return s.Results
}

func (s *BatchDeleteMetaEntitiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteMetaEntitiesResponseBody) SetRequestId(v string) *BatchDeleteMetaEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteMetaEntitiesResponseBody) SetResults(v []*MetaEntityWriteResult) *BatchDeleteMetaEntitiesResponseBody {
	s.Results = v
	return s
}

func (s *BatchDeleteMetaEntitiesResponseBody) SetSuccess(v bool) *BatchDeleteMetaEntitiesResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteMetaEntitiesResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
