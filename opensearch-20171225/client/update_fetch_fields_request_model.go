// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFetchFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []*string) *UpdateFetchFieldsRequest
	GetBody() []*string
	SetDryRun(v bool) *UpdateFetchFieldsRequest
	GetDryRun() *bool
}

type UpdateFetchFieldsRequest struct {
	// The request body.
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateFetchFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFetchFieldsRequest) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsRequest) GetBody() []*string {
	return s.Body
}

func (s *UpdateFetchFieldsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateFetchFieldsRequest) SetBody(v []*string) *UpdateFetchFieldsRequest {
	s.Body = v
	return s
}

func (s *UpdateFetchFieldsRequest) SetDryRun(v bool) *UpdateFetchFieldsRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateFetchFieldsRequest) Validate() error {
	return dara.Validate(s)
}
