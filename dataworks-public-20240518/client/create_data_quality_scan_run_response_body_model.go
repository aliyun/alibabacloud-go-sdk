// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDataQualityScanRunResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDataQualityScanRunResponseBody
	GetRequestId() *string
}

type CreateDataQualityScanRunResponseBody struct {
	// The RunId that was successfully triggered.
	//
	// example:
	//
	// 248840
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityScanRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanRunResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanRunResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDataQualityScanRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityScanRunResponseBody) SetId(v int64) *CreateDataQualityScanRunResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataQualityScanRunResponseBody) SetRequestId(v string) *CreateDataQualityScanRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityScanRunResponseBody) Validate() error {
	return dara.Validate(s)
}
