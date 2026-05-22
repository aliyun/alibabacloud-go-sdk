// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPerformanceDataCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *GetPerformanceDataCollectionResponseBody
	GetEnable() *string
	SetRequestId(v string) *GetPerformanceDataCollectionResponseBody
	GetRequestId() *string
}

type GetPerformanceDataCollectionResponseBody struct {
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPerformanceDataCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPerformanceDataCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPerformanceDataCollectionResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetPerformanceDataCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPerformanceDataCollectionResponseBody) SetEnable(v string) *GetPerformanceDataCollectionResponseBody {
	s.Enable = &v
	return s
}

func (s *GetPerformanceDataCollectionResponseBody) SetRequestId(v string) *GetPerformanceDataCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPerformanceDataCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
