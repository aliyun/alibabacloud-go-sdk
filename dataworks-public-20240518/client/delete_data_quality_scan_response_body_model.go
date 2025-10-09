// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataQualityScanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataQualityScanResponseBody
	GetSuccess() *bool
}

type DeleteDataQualityScanResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2197B9C4-39CE-55EA-8EEA-FDBAE52DXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataQualityScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityScanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataQualityScanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataQualityScanResponseBody) SetRequestId(v string) *DeleteDataQualityScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataQualityScanResponseBody) SetSuccess(v bool) *DeleteDataQualityScanResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataQualityScanResponseBody) Validate() error {
	return dara.Validate(s)
}
