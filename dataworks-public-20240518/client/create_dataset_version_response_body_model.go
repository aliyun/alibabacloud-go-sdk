// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateDatasetVersionResponseBody
	GetId() *string
	SetRequestId(v string) *CreateDatasetVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDatasetVersionResponseBody
	GetSuccess() *bool
}

type CreateDatasetVersionResponseBody struct {
	// ID
	//
	// example:
	//
	// dataworks-datasetVersion:3pXXXb8o0ngr07njhps1
	//
	// :2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDatasetVersionResponseBody) SetId(v string) *CreateDatasetVersionResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDatasetVersionResponseBody) SetRequestId(v string) *CreateDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetVersionResponseBody) SetSuccess(v bool) *CreateDatasetVersionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
