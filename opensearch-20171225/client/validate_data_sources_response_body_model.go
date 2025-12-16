// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ValidateDataSourcesResponseBody
	GetRequestId() *string
	SetResult(v []*ValidateDataSourcesResponseBodyResult) *ValidateDataSourcesResponseBody
	GetResult() []*ValidateDataSourcesResponseBodyResult
}

type ValidateDataSourcesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8FA2B338-AFDC-46B4-A132-B5487820C2BF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	//
	// example:
	//
	// []
	Result []*ValidateDataSourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ValidateDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateDataSourcesResponseBody) GetResult() []*ValidateDataSourcesResponseBodyResult {
	return s.Result
}

func (s *ValidateDataSourcesResponseBody) SetRequestId(v string) *ValidateDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateDataSourcesResponseBody) SetResult(v []*ValidateDataSourcesResponseBodyResult) *ValidateDataSourcesResponseBody {
	s.Result = v
	return s
}

func (s *ValidateDataSourcesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ValidateDataSourcesResponseBodyResult struct {
	// The code returned for the request.
	//
	// example:
	//
	// SUCCEED
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data source.
	//
	// example:
	//
	// {}
	DataSource *ValidateDataSourcesResponseBodyResultDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// The status of the execution.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ValidateDataSourcesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ValidateDataSourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBodyResult) GetCode() *string {
	return s.Code
}

func (s *ValidateDataSourcesResponseBodyResult) GetDataSource() *ValidateDataSourcesResponseBodyResultDataSource {
	return s.DataSource
}

func (s *ValidateDataSourcesResponseBodyResult) GetMessage() *string {
	return s.Message
}

func (s *ValidateDataSourcesResponseBodyResult) SetCode(v string) *ValidateDataSourcesResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResult) SetDataSource(v *ValidateDataSourcesResponseBodyResultDataSource) *ValidateDataSourcesResponseBodyResult {
	s.DataSource = v
	return s
}

func (s *ValidateDataSourcesResponseBodyResult) SetMessage(v string) *ValidateDataSourcesResponseBodyResult {
	s.Message = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResult) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateDataSourcesResponseBodyResultDataSource struct {
	// The parameters of the data source.
	//
	// example:
	//
	// {}
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// user_activity_decision
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// rds
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ValidateDataSourcesResponseBodyResultDataSource) String() string {
	return dara.Prettify(s)
}

func (s ValidateDataSourcesResponseBodyResultDataSource) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) GetTableName() *string {
	return s.TableName
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) GetType() *string {
	return s.Type
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetParameters(v map[string]interface{}) *ValidateDataSourcesResponseBodyResultDataSource {
	s.Parameters = v
	return s
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetTableName(v string) *ValidateDataSourcesResponseBodyResultDataSource {
	s.TableName = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetType(v string) *ValidateDataSourcesResponseBodyResultDataSource {
	s.Type = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) Validate() error {
	return dara.Validate(s)
}
