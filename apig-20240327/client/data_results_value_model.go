// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataResultsValue interface {
	dara.Model
	String() string
	GoString() string
	SetRouteId(v string) *DataResultsValue
	GetRouteId() *string
	SetStatus(v string) *DataResultsValue
	GetStatus() *string
	SetErrorMessage(v string) *DataResultsValue
	GetErrorMessage() *string
	SetSkipReason(v string) *DataResultsValue
	GetSkipReason() *string
}

type DataResultsValue struct {
	RouteId      *string `json:"routeId,omitempty" xml:"routeId,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	SkipReason   *string `json:"skipReason,omitempty" xml:"skipReason,omitempty"`
}

func (s DataResultsValue) String() string {
	return dara.Prettify(s)
}

func (s DataResultsValue) GoString() string {
	return s.String()
}

func (s *DataResultsValue) GetRouteId() *string {
	return s.RouteId
}

func (s *DataResultsValue) GetStatus() *string {
	return s.Status
}

func (s *DataResultsValue) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DataResultsValue) GetSkipReason() *string {
	return s.SkipReason
}

func (s *DataResultsValue) SetRouteId(v string) *DataResultsValue {
	s.RouteId = &v
	return s
}

func (s *DataResultsValue) SetStatus(v string) *DataResultsValue {
	s.Status = &v
	return s
}

func (s *DataResultsValue) SetErrorMessage(v string) *DataResultsValue {
	s.ErrorMessage = &v
	return s
}

func (s *DataResultsValue) SetSkipReason(v string) *DataResultsValue {
	s.SkipReason = &v
	return s
}

func (s *DataResultsValue) Validate() error {
	return dara.Validate(s)
}
