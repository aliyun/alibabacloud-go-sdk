// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTestResultListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v string) *GetTestResultListRequest
	GetConditions() *string
	SetDirectoryIdentifier(v string) *GetTestResultListRequest
	GetDirectoryIdentifier() *string
}

type GetTestResultListRequest struct {
	// example:
	//
	// {\\"conditionGroups\\": [[{\\"fieldIdentifier\\": \\"gmtModified\\", \\"operator\\": \\"MORE_THAN_AND_EQUAL\\", \\"value\\": [\\"2023-04-20 18:03:12.442140\\"], \\"className\\": \\"dateTime\\", \\"format\\": \\"input\\"}]]}
	Conditions *string `json:"conditions,omitempty" xml:"conditions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e27b8eace6501ce51cf5d56784
	DirectoryIdentifier *string `json:"directoryIdentifier,omitempty" xml:"directoryIdentifier,omitempty"`
}

func (s GetTestResultListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTestResultListRequest) GoString() string {
	return s.String()
}

func (s *GetTestResultListRequest) GetConditions() *string {
	return s.Conditions
}

func (s *GetTestResultListRequest) GetDirectoryIdentifier() *string {
	return s.DirectoryIdentifier
}

func (s *GetTestResultListRequest) SetConditions(v string) *GetTestResultListRequest {
	s.Conditions = &v
	return s
}

func (s *GetTestResultListRequest) SetDirectoryIdentifier(v string) *GetTestResultListRequest {
	s.DirectoryIdentifier = &v
	return s
}

func (s *GetTestResultListRequest) Validate() error {
	return dara.Validate(s)
}
