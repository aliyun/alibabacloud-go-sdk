// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateMetaEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchCreateMetaEntitiesResponseBody
	GetRequestId() *string
	SetResults(v []*MetaEntityWriteResult) *BatchCreateMetaEntitiesResponseBody
	GetResults() []*MetaEntityWriteResult
	SetSuccess(v bool) *BatchCreateMetaEntitiesResponseBody
	GetSuccess() *bool
}

type BatchCreateMetaEntitiesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9E0C8E7A-C6BE-5A73-9562-2A030A80E8C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of results for the batch creation operation. This list provides the success status and any error messages for each individual entity in the request.
	Results []*MetaEntityWriteResult `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. This parameter returns `true` even if creating some entities fails. To determine the outcome for each entity, check the `Success` and `ErrorMessage` fields in the `Results` array.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreateMetaEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMetaEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateMetaEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchCreateMetaEntitiesResponseBody) GetResults() []*MetaEntityWriteResult {
	return s.Results
}

func (s *BatchCreateMetaEntitiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchCreateMetaEntitiesResponseBody) SetRequestId(v string) *BatchCreateMetaEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateMetaEntitiesResponseBody) SetResults(v []*MetaEntityWriteResult) *BatchCreateMetaEntitiesResponseBody {
	s.Results = v
	return s
}

func (s *BatchCreateMetaEntitiesResponseBody) SetSuccess(v bool) *BatchCreateMetaEntitiesResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateMetaEntitiesResponseBody) Validate() error {
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
