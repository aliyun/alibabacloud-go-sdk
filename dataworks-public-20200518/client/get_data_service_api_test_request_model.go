// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceApiTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTestId(v int64) *GetDataServiceApiTestRequest
	GetTestId() *int64
}

type GetDataServiceApiTestRequest struct {
	// The Id of the test. TestDataServiceApi is executed asynchronously after the API is called and the test Id is returned. You can also use ListDataServiceApiTest to obtain the latest test Id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123434
	TestId *int64 `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s GetDataServiceApiTestRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceApiTestRequest) GoString() string {
	return s.String()
}

func (s *GetDataServiceApiTestRequest) GetTestId() *int64 {
	return s.TestId
}

func (s *GetDataServiceApiTestRequest) SetTestId(v int64) *GetDataServiceApiTestRequest {
	s.TestId = &v
	return s
}

func (s *GetDataServiceApiTestRequest) Validate() error {
	return dara.Validate(s)
}
