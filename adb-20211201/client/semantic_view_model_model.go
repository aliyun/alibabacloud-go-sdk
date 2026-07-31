// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSemanticViewModel interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *SemanticViewModel
	GetComment() *string
	SetDefinition(v string) *SemanticViewModel
	GetDefinition() *string
	SetScore(v float64) *SemanticViewModel
	GetScore() *float64
	SetViewName(v string) *SemanticViewModel
	GetViewName() *string
	SetViewSchema(v string) *SemanticViewModel
	GetViewSchema() *string
}

type SemanticViewModel struct {
	// The annotation for the semantic view
	//
	// example:
	//
	// 这是一个定义销售额相关指标的视图
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The YAML definition of the semantic view
	//
	// example:
	//
	// name: revenue_analysis
	//
	// description: "Semantic view for analyzing revenue across products and customers"
	//
	// tables:
	//
	//   - name: customers
	//
	//     description: "Customer information"
	//
	//     base_table:
	//
	//       schema: sales_db
	//
	//       table: customers
	//
	//     dimensions:
	//
	//       - name: customer_name
	//
	//         synonyms: ["client name", "customer"]
	//
	//         description: "Full name of the customer"
	//
	//         expr: c_name
	//
	//         data_type: VARCHAR
	//
	//       - name: customer_segment
	//
	//         synonyms: ["segment", "market segment"]
	//
	//         description: "Customer market segment"
	//
	//         expr: c_mktsegment
	//
	//         data_type: VARCHAR
	//
	//         is_enum: true
	//
	//   - name: orders
	//
	//     description: "Order information"
	//
	//     base_table:
	//
	//       schema: sales_db
	//
	//       table: orders
	//
	//     dimensions:
	//
	//       - name: order_date
	//
	//         description: "Date when order was placed"
	//
	//         expr: o_orderdate
	//
	//         data_type: DATE
	//
	//       - name: order_year
	//
	//         description: "Year when order was placed"
	//
	//         expr: YEAR(o_orderdate)
	//
	//         data_type: NUMBER
	//
	//     facts:
	//
	//       - name: order_total
	//
	//         description: "Total order amount"
	//
	//         expr: o_totalprice
	//
	//         data_type: NUMBER
	//
	//     metrics:
	//
	//       - name: total_orders
	//
	//         description: "Total number of orders"
	//
	//         expr: COUNT(*)
	//
	//       - name: total_revenue
	//
	//         description: "Total revenue of orders"
	//
	//         expr: SUM(o_totalprice)
	//
	//       - name: average_order_value
	//
	//         description: "Average order value"
	//
	//         expr: AVG(o_totalprice)
	//
	// relationships:
	//
	//   - name: orders_to_customers
	//
	//     left_table: orders
	//
	//     right_table: customers
	//
	//     relationship_columns:
	//
	//       - left_column: o_custkey
	//
	//         right_column: c_custkey
	//
	// metrics:
	//
	//   - name: revenue_per_customer
	//
	//     description: "Average revenue per customer"
	//
	//     expr: orders.total_revenue / customers.customer_count
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The vector retrieval match score (defaults to 1; during retrieval queries, it is a decimal between 0 and 1 representing vector similarity)
	//
	// example:
	//
	// 0.81
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The name of the semantic view
	//
	// example:
	//
	// revenue_analysis
	ViewName *string `json:"ViewName,omitempty" xml:"ViewName,omitempty"`
	// The schema where the semantic view resides
	//
	// example:
	//
	// sales_db
	ViewSchema *string `json:"ViewSchema,omitempty" xml:"ViewSchema,omitempty"`
}

func (s SemanticViewModel) String() string {
	return dara.Prettify(s)
}

func (s SemanticViewModel) GoString() string {
	return s.String()
}

func (s *SemanticViewModel) GetComment() *string {
	return s.Comment
}

func (s *SemanticViewModel) GetDefinition() *string {
	return s.Definition
}

func (s *SemanticViewModel) GetScore() *float64 {
	return s.Score
}

func (s *SemanticViewModel) GetViewName() *string {
	return s.ViewName
}

func (s *SemanticViewModel) GetViewSchema() *string {
	return s.ViewSchema
}

func (s *SemanticViewModel) SetComment(v string) *SemanticViewModel {
	s.Comment = &v
	return s
}

func (s *SemanticViewModel) SetDefinition(v string) *SemanticViewModel {
	s.Definition = &v
	return s
}

func (s *SemanticViewModel) SetScore(v float64) *SemanticViewModel {
	s.Score = &v
	return s
}

func (s *SemanticViewModel) SetViewName(v string) *SemanticViewModel {
	s.ViewName = &v
	return s
}

func (s *SemanticViewModel) SetViewSchema(v string) *SemanticViewModel {
	s.ViewSchema = &v
	return s
}

func (s *SemanticViewModel) Validate() error {
	return dara.Validate(s)
}
